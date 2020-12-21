package main

import (
	"github.com/ivansukach/linear-congruential-and-MacLaren-Marsaglia-generators/generators"
	"github.com/sirupsen/logrus"
	"math"
)

func main() {
	a01 := 296454621
	a02 := 302711857
	c1 := 48840859
	c2 := 37330745
	M := int(math.Pow(2, 31))
	K := 64
	n := 1000
	aSequence1 := *generators.LinearCongruential(a01, c1, M, n)
	logrus.Info("M: ", M)
	logrus.Info("First element of sequence 1: ", aSequence1[0])
	logrus.Info("Latest element of sequence 1: ", aSequence1[n-1])
	aSequence2 := *generators.LinearCongruential(a02, c2, M, n)
	logrus.Info("First element of sequence 2: ", aSequence2[0])
	logrus.Info("Latest element of sequence 2: ", aSequence2[n-1])
	aSequence1SpecialSize := *generators.LinearCongruential(a01, c1, M, n+K)
	sequenceByMacLarenMarsaglia := *generators.MacLarenMarsaglia(aSequence1SpecialSize, aSequence2, K, n)
	logrus.Info("First element of sequence by MacLarenMarsaglia: ", sequenceByMacLarenMarsaglia[0])
	logrus.Info("Latest element of sequence by MacLarenMarsaglia: ", sequenceByMacLarenMarsaglia[n-1])
}
