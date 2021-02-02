import express, { Express } from 'express';
import { urlencoded } from 'body-parser';

import router from './router';
import errorHandler from '@middleWares/errorHandler';

const app: Express = express();

app.use(urlencoded({ extended: true })).
    use((req, res, next) => {
        res.append('Access-Control-Allow-Origin', '*').
            append('Access-Control-Allow-Methods', 'GET, PUT, POST, DELETE').
            append('Access-Control-Allow-Headers', 'Content-Type, Authorization').
            append('Access-Control-Allow-Credentials', 'true');
        next();
    });


app.use('/api', router)
   .use(errorHandler);

export default app;
