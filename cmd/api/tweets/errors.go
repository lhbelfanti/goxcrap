package tweets

import "errors"

var (
	FailedToRetrieveArticles       = errors.New("failed to retrieve articles")
	EmptyStateNoArticlesToRetrieve = errors.New("empty state, no articles to retrieve")

	FailedToObtainTweetAuthorElement       = errors.New("failed to obtain tweet author element")
	FailedToObtainTweetAuthor              = errors.New("failed to obtain tweet author")
	FailedToObtainQuotedTweetAuthorElement = errors.New("failed to obtain quoted tweet author element")
	FailedToObtainQuotedTweetAuthor        = errors.New("failed to obtain quoted tweet author")
	FailedToObtainTweetAuthorInformation   = errors.New("failed to obtain tweet author information")

	FailedToObtainTweetAvatarElement            = errors.New("failed to obtain tweet avatar element")
	FailedToObtainTweetAvatarImage              = errors.New("failed to obtain tweet avatar image")
	FailedToObtainTweetAvatarSrcFromImage       = errors.New("failed to obtain tweet avatar src from image")
	FailedToObtainQuotedTweetAvatarElement      = errors.New("failed to obtain quoted tweet avatar element")
	FailedToObtainQuotedTweetAvatarImage        = errors.New("failed to obtain quoted tweet avatar image")
	FailedToObtainQuotedTweetAvatarSrcFromImage = errors.New("failed to obtain quoted tweet avatar src from image")

	FailedToObtainTweetIDElement  = errors.New("failed to obtain tweet ID element")
	FailedToObtainTweetIDATag     = errors.New("failed to obtain tweet ID a tag")
	FailedToObtainTweetIDATagHref = errors.New("failed to obtain tweet id a tag href")

	FailedToObtainTweetTimestampElement       = errors.New("failed to obtain tweet timestamp element")
	FailedToObtainTweetTimestampTimeTag       = errors.New("failed to obtain tweet timestamp time tag")
	FailedToObtainTweetTimestamp              = errors.New("failed to obtain tweet timestamp")
	FailedToObtainQuotedTweetTimestampElement = errors.New("failed to obtain quoted tweet timestamp element")
	FailedToObtainQuotedTweetTimestampTimeTag = errors.New("failed to obtain quoted tweet timestamp time tag")
	FailedToObtainQuotedTweetTimestamp        = errors.New("failed to obtain quoted tweet timestamp")
	FailedToObtainTweetTimestampInformation   = errors.New("failed to obtain tweet timestamp information")

	FailedToObtainTweetTextElement           = errors.New("failed to obtain tweet text element")
	FailedToObtainTweetTextParts             = errors.New("failed to obtain tweet text parts")
	FailedToObtainTweetTextPartTagName       = errors.New("failed to obtain tweet text part tag name")
	FailedToObtainTweetTextFromSpan          = errors.New("failed to obtain tweet text from span")
	FailedToObtainQuotedTweetTextElement     = errors.New("failed to obtain quoted tweet text element")
	FailedToObtainQuotedTweetTextParts       = errors.New("failed to obtain quoted tweet text parts")
	FailedToObtainQuotedTweetTextPartTagName = errors.New("failed to obtain quoted tweet text part tag name")
	FailedToObtainQuotedTweetTextFromSpan    = errors.New("failed to obtain quoted tweet text from span")

	FailedToObtainTweetImagesElement       = errors.New("failed to obtain tweet images element")
	FailedToObtainTweetImages              = errors.New("failed to obtain tweet images")
	FailedToObtainTweetSrcFromImage        = errors.New("failed to obtain tweet src from image")
	FailedToObtainQuotedTweetImagesElement = errors.New("failed to obtain quoted tweet images element")
	FailedToObtainQuotedTweetImages        = errors.New("failed to obtain quoted tweet images")
	FailedToObtainQuotedTweetSrcFromImage  = errors.New("failed to obtain quoted tweet src from image")
)
