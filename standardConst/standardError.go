package standardConst

import "fmt"

var ChallengeError = fmt.Errorf("挑战值错误，请稍后")

const (
	CycleStageGen      = iota //gen proof
	CycleStageCollect         // collect minerals
	CycleStageAnnounce        // announce winner and check spv proof
	CycleStageUpdate          // update proof
	CycleStageWait            // wait a moment
)

const (
	UrlUpdateOrGen = "/updateOrGen"
	UrlMining      = "/mining"
	UrlProve       = "/prove"
)
