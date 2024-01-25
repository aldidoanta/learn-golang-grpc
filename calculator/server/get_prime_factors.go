package main

import (
	"log"

	pb "github.com/aldidoanta/learn-golang-grpc/calculator/proto"
)

func (s Server) GetPrimeFactors(in *pb.PrimeRequest, stream pb.CalculatorService_GetPrimeFactorsServer) error {
	log.Printf("GetPrimeFactors() was invoked with %v\n", in)

	// the algorithm of prime factorization
	var divisor int32 = 2
	var number int32 = in.Number

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				PrimeFactor: divisor,
			})

			log.Println(divisor)
			number = number / divisor
		} else {
			divisor = divisor + 1
		}

	}

	return nil
}
