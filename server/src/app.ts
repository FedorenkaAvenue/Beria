import express, { Express } from 'express';
import { urlencoded, json } from 'body-parser';

import errorHandler from '@middleWares/errorHandler';

const app: Express = express();

app.use(urlencoded({ extended: true })).
    // use(json()).
    use((req, res, next) => {
        res.append('Access-Control-Allow-Origin', '*').
            append('Access-Control-Allow-Methods', 'GET, PUT, POST, DELETE').
            append('Access-Control-Allow-Headers', 'Content-Type, Authorization').
            append('Access-Control-Allow-Credentials', 'true');
        next();
    });

app.get('/', (req, res) => res.json('Beria in da house!'))

app.use(errorHandler);

export default app;
