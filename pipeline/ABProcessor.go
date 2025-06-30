package main

import "context"

type aOut struct {
}

type bOut struct {
}

type cIn struct {
	a aOut
	b bOut
}

type Input struct {
	A string
	B string
}

type ABProcessor struct {
	outA chan aOut
	outB chan bOut
	errs chan error
}

func NewABProcessor() *ABProcessor {
	return &ABProcessor{
		outA: make(chan aOut, 1),
		outB: make(chan bOut, 1),
		errs: make(chan error, 2),
	}
}

func (p *ABProcessor) start(ctx context.Context, data Input) {
	go func() {
		res, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- res
	}()

	go func() {
		res, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- res
	}()
}

func (p *ABProcessor) wait(ctx context.Context) (cIn, error) {
	var cData cIn
	for count := 0; count < 2; count++ {
		select {
		case a := <-p.outA:
			cData.a = a
		case b := <-p.outB:
			cData.b = b
		case err := <-p.errs:
			return cIn{}, err
		case <-ctx.Done():
			return cIn{}, ctx.Err()
		}
	}
	return cData, nil
}

func getResultA(ctx context.Context, in string) (aOut, error) {
	return aOut{}, nil
}

func getResultB(ctx context.Context, in string) (bOut, error) {
	return bOut{}, nil
}
