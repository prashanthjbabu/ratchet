package util

import "github.com/dailyburn/ratchet/logger"

// KillPipelineIfErr is an error-checking helper.
func KillPipelineIfErr(err error, killChan chan error) {
	if err != nil {
		logger.Error(err.Error())
		/*Make this a non-blocking call incase killchan was signalled earlier through some other thread*/
		select {
		case killChan <- err:
		default:
		}
	}
}
