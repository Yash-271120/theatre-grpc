package grpcserver

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	grpcclient "github.com/Yash-271120/theatre-grpc/backend/services/ticket/grpcClient"
	"github.com/Yash-271120/theatre-grpc/backend/services/ticket/models"
	pb "github.com/Yash-271120/theatre-grpc/backend/services/ticket/proto"
)

type TicketRPCServer struct {
	pb.TicketServiceServer
}

func (s *TicketRPCServer) CreateTicket(req *pb.CreateTicketRequest, stream pb.TicketService_CreateTicketServer) error {
	log.Println("CreateTicket called")
	hall_id := req.GetHallId()
	seats := req.GetSeats()

	response := pb.CreateTicketResponse{
		Status: pb.TicketBookingStatus_TICKET_BOOKING_STATUS_PENDING,
	}

	if err := stream.Send(&response); err != nil {
		log.Fatalf("Error while sending response : %v", err)
		return err
	}

	// call hall service to update hall
	update_hall_req := &pb.UpdateHallAvailableRequest{
		Id:        hall_id,
		Available: seats,
	}

	time.Sleep(1000 * time.Millisecond)

	response.Status = pb.TicketBookingStatus_TICKET_BOOKING_STATUS_RESERVING_SEATS

	if err := stream.Send(&response); err != nil {
		log.Fatalf("Error while sending response : %v", err)
		return err
	}

	_, err := grpcclient.CallUpdateHall(update_hall_req, &grpcclient.HallRpcServiceObj)

	if err != nil {
		log.Fatalf("Error while calling UpdateHall RPC : %v", err)
		return err
	}

	time.Sleep(3000 * time.Millisecond)
	ticket := models.Ticket{
		HallId:    uint32(hall_id),
		Seats:     uint32(seats),
		EntryCode: generateEntryCode(),
	}
	models.DB.Create(&ticket)
	response.Ticket = &pb.Ticket{
		Id:        int32(ticket.Id),
		HallId:    int32(ticket.HallId),
		Seats:     int32(ticket.Seats),
		EntryCode: ticket.EntryCode,
	}

	response.Status = pb.TicketBookingStatus_TICKET_BOOKING_STATUS_CONFIRMED

	if err := stream.Send(&response); err != nil {
		log.Fatalf("Error while sending response : %v", err)
		return err
	}
	return nil

}

func generateEntryCode() string {
	// Generate 2 random numbers
	numbers := ""
	for i := 0; i < 2; i++ {
		numbers += fmt.Sprint((rand.Intn(10))) // ASCII values for digits 0-9
	}

	// Generate 4 random capital letters
	letters := ""
	for i := 0; i < 4; i++ {
		letters += string((rand.Intn(26) + 65)) // ASCII values for capital letters A-Z
	}

	// Combine the numbers and letters to create the entry code
	entryCode := numbers + letters

	return entryCode
}
