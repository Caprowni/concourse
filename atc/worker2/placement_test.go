package worker2_test

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/atc/worker2"
	grt "github.com/concourse/concourse/atc/worker2/gardenruntime/gardenruntimetest"
	"github.com/concourse/concourse/atc/worker2/workertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = Describe("Container Placement Strategies", func() {
	Describe("Volume Locality", func() {
		volumeLocalityStrategy := func() worker2.PlacementStrategy {
			strategy, err := worker2.NewPlacementStrategy(worker2.PlacementOptions{
				Strategies: []string{"volume-locality"},
			})
			Expect(err).ToNot(HaveOccurred())
			return strategy
		}

		Test("selects the worker with the most inputs locally", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input1"),
							grt.NewVolume("input3"),
						),
					grt.NewWorker("worker2").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input2"),
						),
				),
			)

			worker, err := volumeLocalityStrategy().Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,

				Inputs: []runtime.Input{
					{
						VolumeHandle:    "input1",
						DestinationPath: "/input1",
					},
					{
						VolumeHandle:    "input2",
						DestinationPath: "/input2",
					},
					{
						VolumeHandle:    "input3",
						DestinationPath: "/input3",
					},
				},
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(Equal("worker1"))
		})

		Test("includes all workers in the case of a tie", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input1"),
						),
					grt.NewWorker("worker2").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input2"),
						),
				),
			)

			worker, err := volumeLocalityStrategy().Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,

				Inputs: []runtime.Input{
					{
						VolumeHandle:    "input1",
						DestinationPath: "/input1",
					},
					{
						VolumeHandle:    "input2",
						DestinationPath: "/input2",
					},
				},
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(BeOneOf("worker1", "worker2"))
		})

		Test("considers resource caches", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithDBContainersInState(grt.Creating, "container1").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input1"),
							grt.NewVolume("cache-input2"),
						).
						WithResourceCacheOnVolume("container1", "cache-input2", "some-resource"),
					grt.NewWorker("worker2").
						WithDBContainersInState(grt.Creating, "container2").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input2"),
						).
						WithResourceCacheOnVolume("container2", "input2", "some-resource"),
				),
			)

			worker, err := volumeLocalityStrategy().Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,

				Inputs: []runtime.Input{
					{
						VolumeHandle:    "input1",
						DestinationPath: "/input1",
					},
					{
						VolumeHandle:    "input2",
						DestinationPath: "/input2",
					},
				},
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(Equal("worker1"))
		})

		Test("considers task caches", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input1"),
						).
						WithCachedPaths("/cache1", "/cache2"),
					grt.NewWorker("worker2").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input2"),
						).
						WithCachedPaths("/cache1"),
				),
			)

			worker, err := volumeLocalityStrategy().Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,

				Inputs: []runtime.Input{
					{
						VolumeHandle:    "input1",
						DestinationPath: "/input1",
					},
					{
						VolumeHandle:    "input2",
						DestinationPath: "/input2",
					},
				},

				Caches: []string{"/cache1", "/cache2"},
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(Equal("worker1"))
		})

		Test("does not consider workers that have been filtered out", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1"),
					grt.NewWorker("worker2").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("input1"),
						),
				),
			)

			worker, err := volumeLocalityStrategy().Choose(
				logger,
				scenario.Pool,
				filterWorkers(scenario.DB.Workers, "worker1"),
				runtime.ContainerSpec{
					TeamID:   scenario.TeamID,
					JobID:    scenario.JobID,
					StepName: scenario.StepName,

					Inputs: []runtime.Input{
						{
							VolumeHandle:    "input1",
							DestinationPath: "/input1",
						},
					},
				},
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(Equal("worker1"))
		})
	})

	Describe("Fewest Build Containers", func() {
		fewestBuildContainersStrategy := func() worker2.PlacementStrategy {
			strategy, err := worker2.NewPlacementStrategy(worker2.PlacementOptions{
				Strategies: []string{"fewest-build-containers"},
			})
			Expect(err).ToNot(HaveOccurred())
			return strategy
		}

		Test("returns workers with the fewest active containers", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c1"),
						),
					grt.NewWorker("worker2").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c2"),
							grt.NewContainer("c3"),
						),
					grt.NewWorker("worker3").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c4"),
						),
				),
			)

			worker, err := fewestBuildContainersStrategy().Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(BeOneOf("worker1", "worker3"))
		})
	})

	Describe("Limit Active Tasks", func() {
		limitActiveTasksStrategy := func(max int) worker2.PlacementStrategy {
			strategy, err := worker2.NewPlacementStrategy(worker2.PlacementOptions{
				Strategies:              []string{"limit-active-tasks"},
				MaxActiveTasksPerWorker: max,
			})
			Expect(err).ToNot(HaveOccurred())
			return strategy
		}

		Test("returns workers with the fewest active tasks", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithActiveTasks(1),
					grt.NewWorker("worker2").
						WithActiveTasks(2),
					grt.NewWorker("worker3").
						WithActiveTasks(1),
				),
			)

			worker, err := limitActiveTasksStrategy(0).Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(BeOneOf("worker1", "worker3"))
		})

		Test("allows setting a limit on the number of active tasks", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithActiveTasks(10),
					grt.NewWorker("worker2").
						WithActiveTasks(20),
					grt.NewWorker("worker3").
						WithActiveTasks(10),
				),
			)

			_, err := limitActiveTasksStrategy(10).Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,

				Type: db.ContainerTypeTask,
			})
			Expect(err).To(MatchError("no worker fit container placement strategy: limit-active-tasks"))
		})

		Test("limit only applies to task containers", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithActiveTasks(10),
					grt.NewWorker("worker2").
						WithActiveTasks(20),
					grt.NewWorker("worker3").
						WithActiveTasks(10),
				),
			)

			worker, err := limitActiveTasksStrategy(10).Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,

				Type: db.ContainerTypeCheck,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(BeOneOf("worker1", "worker3"))
		})
	})

	Describe("Limit Active Containers", func() {
		limitActiveContainersStrategy := func(max int) worker2.PlacementStrategy {
			strategy, err := worker2.NewPlacementStrategy(worker2.PlacementOptions{
				Strategies:                   []string{"limit-active-containers"},
				MaxActiveContainersPerWorker: max,
			})
			Expect(err).ToNot(HaveOccurred())
			return strategy
		}

		Test("removes workers with too many active containers", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c1"),
						),
					grt.NewWorker("worker2").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c2"),
							grt.NewContainer("c3"),
						),
					grt.NewWorker("worker3").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c4"),
						),
				),
			)

			worker, err := limitActiveContainersStrategy(2).Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(BeOneOf("worker1", "worker3"))
		})

		Test("noop if limit is unset", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c1"),
						),
					grt.NewWorker("worker2").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c2"),
							grt.NewContainer("c3"),
						),
					grt.NewWorker("worker3").
						WithContainersCreatedInDBAndGarden(
							grt.NewContainer("c4"),
						),
				),
			)

			worker, err := limitActiveContainersStrategy(0).Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(BeOneOf("worker1", "worker2", "worker3"))
		})
	})

	Describe("Limit Active Volumes", func() {
		limitActiveVolumesStrategy := func(max int) worker2.PlacementStrategy {
			strategy, err := worker2.NewPlacementStrategy(worker2.PlacementOptions{
				Strategies:                []string{"limit-active-volumes"},
				MaxActiveVolumesPerWorker: max,
			})
			Expect(err).ToNot(HaveOccurred())
			return strategy
		}

		Test("removes workers with too many active volumes", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("v1"),
						),
					grt.NewWorker("worker2").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("v2"),
							grt.NewVolume("v3"),
						),
					grt.NewWorker("worker3").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("v4"),
						),
				),
			)

			worker, err := limitActiveVolumesStrategy(2).Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(BeOneOf("worker1", "worker3"))
		})

		Test("noop if limit is unset", func() {
			scenario := Setup(
				workertest.WithBasicJob(),
				workertest.WithWorkers(
					grt.NewWorker("worker1").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("v1"),
						),
					grt.NewWorker("worker2").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("v2"),
							grt.NewVolume("v3"),
						),
					grt.NewWorker("worker3").
						WithVolumesCreatedInDBAndBaggageclaim(
							grt.NewVolume("v4"),
						),
				),
			)

			worker, err := limitActiveVolumesStrategy(0).Choose(logger, scenario.Pool, scenario.DB.Workers, runtime.ContainerSpec{
				TeamID:   scenario.TeamID,
				JobID:    scenario.JobID,
				StepName: scenario.StepName,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Name()).To(BeOneOf("worker1", "worker2", "worker3"))
		})
	})
})

func BeOneOf(vals ...interface{}) types.GomegaMatcher {
	matchers := make([]types.GomegaMatcher, len(vals))
	for i, v := range vals {
		matchers[i] = Equal(v)
	}
	return SatisfyAny(matchers...)
}

func filterWorkers(allWorkers []db.Worker, namesToKeep ...string) []db.Worker {
	keep := func(name string) bool {
		for _, otherName := range namesToKeep {
			if name == otherName {
				return true
			}
		}
		return false
	}

	var workers []db.Worker
	for _, worker := range allWorkers {
		if keep(worker.Name()) {
			workers = append(workers, worker)
		}
	}
	return workers
}
