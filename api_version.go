// Package egobot provides a complete Go library for creating Telegram bots.
//
// This library is synced with Telegram Bot API version 10.2 (July 14, 2026).
//
// For more information, see: https://core.telegram.org/bots/api
package egobot

// APIVersion returns the Telegram Bot API version that this library is synced with.
func APIVersion() string {
	return "10.2"
}

// APIDate returns the release date of the Telegram Bot API version that this library is synced with.
func APIDate() string {
	return "2026-07-14"
}
