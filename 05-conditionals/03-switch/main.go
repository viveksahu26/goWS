package main

func Direction(d string) string {
	switch d {
	case "N":
		return "North"
	case "E":
		return "East"
	case "S":
		return "South"
	case "W":
		return "West"
	default:
		return "Unknown"
	}
}
