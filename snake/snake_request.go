/*
 * Battlesnake Snake API
 *
 * Battlesnake participants implement this webhook API to power their snake. See battlesnake.io for details.
 *
 * API version: 2018.03.beta
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package snakemodel

type SnakeRequest struct {
	Game Game `json:"game"`
	Turn int32 `json:"turn"`
	Board Board `json:"board"`
	You Snake `json:"you"`
}
