package main

import "context"

type COut struct {
	count map[rune]int
}

type CProcessor struct {
	outC chan COut
	errs chan error
}

func newCProcessor() *CProcessor {
	return &CProcessor{
		outC: make(chan COut, 1),
		errs: make(chan error, 1),
	}
}

func (p *CProcessor) start(ctx context.Context, input cIn) {
	go func() {
		res, err := getResultC(ctx, input)
		if err != nil {
			p.errs <- err
			return
		}
		p.outC <- res
	}()
}

func getResultC(ctx context.Context, input cIn) (COut, error) {
	return COut{}, nil
}

func (p *CProcessor) wait(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}
