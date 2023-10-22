import dotenv from 'dotenv';
import grpc from '@grpc/grpc-js'
import protoLoader from '@grpc/proto-loader'
import path from 'path';

import { GetMovies, CreateMovie } from './controller/movie.js'
import { dbConnect } from './db/index.js';

dotenv.config();
const __dirname = path.resolve();
const protoPath = __dirname + "/../../proto/movieService.proto"

const packageDefination = protoLoader.loadSync(protoPath, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
})

const movieProto = grpc.loadPackageDefinition(packageDefination).proto

const getServer = () => {
    const server = new grpc.Server()
    server.addService(movieProto.MovieService.service, {
        GetMovies,
        CreateMovie
    })
    return server
}

const server = getServer()

server.bindAsync(`${process.env.HOST}:${process.env.PORT}`, grpc.ServerCredentials.createInsecure(), async (err, port) => {
    if (err != null) {
        return console.error(err)
    }
    console.log(`Movie gRPC server listening on ${port}`)
    await dbConnect()
    server.start()
})


