package router

import (
	"log"
	"net/http"

	grpcclient "github.com/Yash-271120/theatre-grpc/backend/server/grpcClient"
	pb "github.com/Yash-271120/theatre-grpc/backend/server/proto"
	response "github.com/Yash-271120/theatre-grpc/backend/server/response"
	"github.com/gin-gonic/gin"
)

type CreateMovieInput struct {
	Title string `json:"title" binding:"required"`
}

func addMovieRoutes(rg *gin.RouterGroup) {
	rg.GET("/movies", func(c *gin.Context) {
		getMovieReq := &pb.GetMoviesRequest{}
		getMovieRes, err := grpcclient.MovieRpcServiceObj.GetMovies(c, getMovieReq)
		if err != nil {
			log.Fatalf("Error when calling GetMovies: %s", err)
			response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SuccessResponse(c, http.StatusOK, getMovieRes.Movies)
	})

	rg.POST("/movie", func(c *gin.Context) {
		var input CreateMovieInput
		if err := c.ShouldBindJSON(&input); err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		createMovieReq := &pb.CreateMovieRequest{
			Title: input.Title,
		}

		createMovieRes, err := grpcclient.MovieRpcServiceObj.CreateMovie(c, createMovieReq)
		if err != nil {
			log.Fatalf("Error when calling CreateMovie: %s", err)
			response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SuccessResponse(c, http.StatusOK, createMovieRes.Movie)
	})
}
