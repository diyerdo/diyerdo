// diyerdo backend implementation
// Copyright (C) 2026 DrLarck
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// ---
//
// diyerdo entrypoint
package main

import (
	"net"

	"github.com/diyerdo/diyerdo/internal/services/equipments/api"
	"github.com/diyerdo/proto/gen/go/proto/equipments/v1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"google.golang.org/grpc"
)

// initLogger is a helper function to initialize the logger
//
// See https://github.com/rs/zerolog
func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

// initListener is a helper function to initialize the TCP listener
//
// On failure it will exit the program, writing what went wrong as Fatal log
// message
func listener() net.Listener {
	listener, err := net.Listen("tcp", ":58180")
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Str("address", listener.Addr().String()).
			Msg("failed to create TCP listener")
	}

	return listener
}

// registerServices is a helper function to register the gRPC services
func registerServicesGrpc(server *grpc.Server) {
	// Equipments service
	equipmentsService := instanciateService(api.NewEquipmentsService)
	equipments.RegisterEquipmentsServiceServer(server, equipmentsService.Server)
}

// instanciateService is a helper function to instanciate a service. It returns
// a pointer to the service instance on success
//
// On failure it will exit the program, writing what went wrong as Fatal log
// message
func instanciateService[T any](constructor func() (*T, error)) *T {
	service, err := constructor()
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msg("failed to create service")
	}

	return service
}

// serve is a helper function to serve the gRPC server
//
// On failure it will exit the program, writing what went wrong as Fatal log
// message
func serve(server *grpc.Server, listener net.Listener) {
	if err := server.Serve(listener); err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Str("address", listener.Addr().String()).
			Msg("failed to serve gRPC server")
	}
}

// main is the entrypoint of the diyerdo backend
func main() {
	initLogger()

	server := grpc.NewServer()
	registerServicesGrpc(server)

	listener := listener()
	log.Info().Msgf("gRPC server listening at %v", listener.Addr().String())
	serve(server, listener)
}
