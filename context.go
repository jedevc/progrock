package progrock

import "context"

func RecorderToContext(ctx context.Context, recorder *Recorder) context.Context {
	return context.WithValue(ctx, recorderKey{}, recorder)
}

func RecorderFromContext(ctx context.Context) *Recorder {
	rec := ctx.Value(recorderKey{})
	if rec == nil {
		return NewRecorder(Discard{})
	}

	return rec.(*Recorder)
}

func WithGroup(ctx context.Context, name string, labels ...*Label) (context.Context, *Recorder) {
	rec := RecorderFromContext(ctx).WithGroup(name, labels...)
	return RecorderToContext(ctx, rec), rec
}

type recorderKey struct{}
