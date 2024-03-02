package src

import "context"

// URLShortenInput ...
type URLShortenInput struct {
	URL URL
}

// URLShortenOutput ...
type URLShortenOutput struct {
	UUID UUID
}

// MakeURLShortenOutput ...
func MakeURLShortenOutput(uuid UUID) Either[URLShortenOutput] {
	return Right(URLShortenOutput{
		UUID: uuid,
	})
}

// URLShortenApplication ...
type URLShortenApplication struct {
	url URLRepository
}

// URLShorten ...
func (a URLShortenApplication) URLShorten(ctx context.Context, input URLShortenInput) Either[URLShortenOutput] {
	return Bind(
		Bind(
			Right(input.URL),
			WithContext(ctx, a.url.Store),
		),
		MakeURLShortenOutput,
	)
}
