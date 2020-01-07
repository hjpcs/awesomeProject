package ch38

import "testing"

func TestCreateRequest(t *testing.T) {
	str := createRequest()
	t.Log(str)
}

func TestProcessRequest(t *testing.T) {
	var reqs []string
	reqs = append(reqs, createRequest())
	reps := processRequest(reqs)
	t.Log(reps[0])
}

func BenchmarkProcessRequest(b *testing.B) {

	var reqs []string
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequest(reqs)
	}
	b.StopTimer()

}

func BenchmarkProcessRequestOld(b *testing.B) {

	var reqs []string
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequestOld(reqs)
	}
	b.StopTimer()

}

//cmdline
//go test -bench=. -cpuprofile=cpu.prof
//top -cum