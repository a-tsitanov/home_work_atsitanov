package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func initChans(countChan int) []Bi {
	channels := make([]Bi, countChan)
	for i := 0; i < countChan; i++ {
		channels[i] = make(Bi)
	}
	return channels
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	countStages := len(stages)
	stageChannels := initChans(countStages)
	doneCoroutine := make(Bi, countStages)

	go func() {
		for {
			<-done
			for i := 0; i < countStages; i++ {
				doneCoroutine <- struct{}{}
			}
		}
	}()

	job := func(stage Stage, in In, out Bi) {
		outStage := stage(in)
		for {
			select {
			case <-doneCoroutine:
				close(out)
				return
			case v, ok := <-outStage:
				if !ok {
					close(out)
					return
				}
				out <- v
			}
		}
	}

	for i, stage := range stages {
		if i == 0 {
			go job(stage, in, stageChannels[i])
		} else {
			go job(stage, stageChannels[i-1], stageChannels[i])
		}
	}

	return stageChannels[countStages-1]
}
