// +build !windows

package updater

// Fork process that inherits listening sockets
// Initialize child and start accepting connections on listening sockets
// Send signal to parent to terminate
func restart() {

}
