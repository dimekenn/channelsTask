package service

type Service interface {
	FromNChannelsToOneChannel(channelCount, randomIn int) []int
}