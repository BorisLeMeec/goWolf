package main

func isThereWall(pos position, myData data) bool {
	if myData.theMap.array[pos.y*myData.theMap.size.x+pos.x] == '1' {
		return true
	}
	return false
}
