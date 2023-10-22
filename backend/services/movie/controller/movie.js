import Movie from '../models/movieModel.js'

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
        callback(null, { movie });
    } catch (error) {
        callback(error, null);
    }
}
