import grpc from '@grpc/grpc-js'
import protoLoader from '@grpc/proto-loader'
import path from 'path';
import dotenv from 'dotenv';

dotenv.config();

const __dirname = path.resolve();
const protoPath = __dirname + "/../../proto/hallService.proto"

const packageDefination = protoLoader.loadSync(protoPath, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
})

const hallProto = grpc.loadPackageDefinition(packageDefination).proto

class HallServiceClient {
    constructor() {
        this.client = new hallProto.HallService(`${process.env.HALL_SERVICE_HOST}:${process.env.HALL_SERVICE_PORT}`, grpc.credentials.createInsecure())
    }
    createHall = (createHallRequest) => {
        const retval = this.client.CreateHall(createHallRequest, (err, response) => {
            if (err) {
                console.log(err)
                throw new Error(err)
            }
            return response
        })
        return retval
    }

    closeConnection = () => {
        this.client.close()
    }
}

export default HallServiceClient