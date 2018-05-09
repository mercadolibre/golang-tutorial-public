type Time struct {
    sec  int64
    nsec int32
    loc  *Location
}

func Now() Time {
  	sec, nsec := now()
  	return Time{sec + unixToInternal, nsec, Local}
}

func (t Time) Add(d Duration) Time {
  	t.sec += int64(d / 1e9)
  	nsec := t.nsec + int32(d%1e9)
  	if nsec >= 1e9 {
  		t.sec++
  		nsec -= 1e9
  	} else if nsec < 0 {
  		t.sec--
  		nsec += 1e9
  	}
  	t.nsec = nsec
  	return t
}
