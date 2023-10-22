import { DataTypes, Sequelize, Model } from 'sequelize';
import sequelize from '../db/index.js';

const Movie = sequelize.define('movie', {
    id: {
        type: DataTypes.INTEGER,
        primaryKey: true,
        autoIncrement: true
    },
    title: {
        type: DataTypes.STRING,
        allowNull: false,
    },
});

export default Movie;

