package ack

// rekursiv funksjon (ackerman, supereksponensielt)
func Ack(m, n int) int {
	var ans int // lokal variabel
	if m == 0 { // hvis null lever tilbake n + 1
		ans = n + 1
	} else if n == 0 {
		ans = Ack(m-1, 1)
	} else {
		ans = Ack(m-1, Ack(m, n-1)) // kall til ackerman er selv en akerman; gjennom mange stabelrammer
	}
	return ans
}
