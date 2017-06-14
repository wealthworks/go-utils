//
// 定时任务，用于资源清理
//
package reaper

//
// Example:
// defer reaper.Quit(reaper.Run(0, func() error {
// 		// Cleanup
// }))
//
