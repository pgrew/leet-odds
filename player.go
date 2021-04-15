package main

type Player struct {
	isActing bool
	isDealer bool
	isMucked bool
	amountPotOwes int64
	numberOfPotSplashes int64
	stack int64 //todo chip
  hand Hand
}
