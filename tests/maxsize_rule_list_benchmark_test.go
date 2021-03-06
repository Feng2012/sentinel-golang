package tests

import (
	"math/rand"
	"strconv"
	"testing"

	cb "github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/util"
)

func Test_Size_1000_Circuit_Breaker_Rules_Update(t *testing.T) {
	rs := make([]cb.Rule, 0, 1000)
	rand.Seed(int64(util.CurrentTimeMillis()))
	intervals := []uint32{10000, 15000, 20000, 25000, 30000}
	for i := 0; i < 1000; i++ {
		retryTimeout := intervals[rand.Int()%5]
		rs = append(rs, cb.NewRule("github.com/alibaba/sentinel/test"+strconv.Itoa(rand.Int()%100), cb.SlowRequestRatio,
			cb.WithStatIntervalMs(10000), cb.WithRetryTimeoutMs(retryTimeout),
			cb.WithMinRequestAmount(rand.Uint64()%100), cb.WithMaxAllowedRtMs(100),
			cb.WithMaxSlowRequestRatio(0.1)))
	}

	_, err := cb.LoadRules(rs)
	if err != nil {
		t.Errorf("error")
	}
}

func Benchmark_Size_1000_Circuit_Breaker_Rules_Update(b *testing.B) {
	rs := make([]cb.Rule, 0, 1000)
	rand.Seed(int64(util.CurrentTimeMillis()))
	intervals := []uint32{10000, 15000, 20000, 25000, 30000}
	for i := 0; i < 1000; i++ {
		retryTimeout := intervals[rand.Int()%5]
		rs = append(rs, cb.NewRule("github.com/alibaba/sentinel/test"+strconv.Itoa(rand.Int()%100), cb.SlowRequestRatio,
			cb.WithStatIntervalMs(10000), cb.WithRetryTimeoutMs(retryTimeout),
			cb.WithMinRequestAmount(rand.Uint64()%100), cb.WithMaxAllowedRtMs(100),
			cb.WithMaxSlowRequestRatio(0.1)))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := cb.LoadRules(rs)
		if err != nil {
			b.Errorf("error")
		}
	}
}

func Test_Size_10000_Circuit_Breaker_Rules_Update(t *testing.T) {
	rs := make([]cb.Rule, 0, 10000)
	intervals := []uint32{10000, 15000, 20000, 25000, 30000}
	for i := 0; i < 10000; i++ {
		retryTimeout := intervals[rand.Int()%5]
		rs = append(rs, cb.NewRule("github.com/alibaba/sentinel/test"+strconv.Itoa(rand.Int()%100), cb.SlowRequestRatio,
			cb.WithStatIntervalMs(10000), cb.WithRetryTimeoutMs(retryTimeout),
			cb.WithMinRequestAmount(rand.Uint64()%100), cb.WithMaxAllowedRtMs(100),
			cb.WithMaxSlowRequestRatio(0.1)))
	}

	_, err := cb.LoadRules(rs)
	if err != nil {
		t.Errorf("error")
	}
}

func Benchmark_Size_10000_Circuit_Breaker_Rules_Update(b *testing.B) {
	rs := make([]cb.Rule, 0, 10000)
	intervals := []uint32{10000, 15000, 20000, 25000, 30000}
	for i := 0; i < 10000; i++ {
		retryTimeout := intervals[rand.Int()%5]
		rs = append(rs, cb.NewRule("github.com/alibaba/sentinel/test"+strconv.Itoa(rand.Int()%100), cb.SlowRequestRatio,
			cb.WithStatIntervalMs(10000), cb.WithRetryTimeoutMs(retryTimeout),
			cb.WithMinRequestAmount(rand.Uint64()%100), cb.WithMaxAllowedRtMs(100),
			cb.WithMaxSlowRequestRatio(0.1)))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := cb.LoadRules(rs)
		if err != nil {
			b.Errorf("error")
		}
	}
}
