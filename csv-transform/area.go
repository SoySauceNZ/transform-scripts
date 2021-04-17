package main

// getCrash in image
func getCrash(im image) []crash {
	var crashes []crash
	for _, cr := range crashesArray {
		if betweenLat(im.latbottom, im.lattop, cr.latitude) &&
			betweenLng(im.lngleft, im.lngright, cr.longitude) {
			cr.image = im.image
			crashes = append(crashes, cr)
		}
	}
	return crashes
}

func betweenLat(lat0, lat1, lat float64) bool {
	if lat0 < lat1 {
		if lat0 <= lat && lat < lat1 {
			return true
		}
	} else if lat1 < lat0 {
		if lat1 <= lat && lat < lat0 {
			return true
		}
	}
	return false
}

func betweenLng(lng0, lng1, lng float64) bool {
	if lng0 < lng1 {
		if lng0 <= lng && lng < lng1 {
			return true
		}
	} else if lng1 < lng0 {
		if lng1 <= lng && lng < lng0 {
			return true
		}
	}
	return false
}
