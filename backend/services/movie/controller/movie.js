import Movie from '../models/movieModel.js'
import HallServiceClient from '../grpcClient/hallServiceClient.js';

export const GetMovies = async (call, callback) => {
    try {
        const movies = await Movie.findAll();
        callback(null, { movies: movies });
    } catch (error) {
        callback(error, null);
    }
}

export const CreateMovie = async (call, callback) => {
    try {
        const { title } = call.request;
        const movie = await Movie.create({ title });
        const newMovieId = movie.id;
        const newMovieTitle = movie.title;
        const hallServiceClient = new HallServiceClient();
        const hallCapacity = Math.floor(Math.random() * 40) + 10;
        const createHallRequest = {
            movie,
            capacity: hallCapacity,
            available: hallCapacity
        }
        hallServiceClient.createHall(createHallRequest, (err, response) => {
            if (err) {
                console.log(err)
                throw new Error(err)
            }
            console.log(response)
        });
        callback(null, { movie });
    } catch (error) {
        callback(error, null);
    }
}
