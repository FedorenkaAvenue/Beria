import { NextFunction, Request, Response } from 'express';

import GET_CITY_LIST from '@db/redis/queries/cities';

type IParams = {
    country: string
}

export default async function(req: Request<IParams>, res: Response, next: NextFunction) {
    const {
        params: { country },
        query: { sub }
    } = req;

    try {
        if (!country) throw Error;

        try {
            const cities = await GET_CITY_LIST(country);

            if (!cities) return res.sendStatus(406);

            if (sub) {
                const subRegExp = new RegExp(`^${sub}`, 'i');

                return res.json(cities.filter(
                    ({ title: { en, origin } }) => subRegExp.test(en) || subRegExp.test(origin)
                ));
            }

            res.json(cities);
        } catch(err) {
            next(err);
        }
    } catch(err) {
        res.sendStatus(400);
    }
}
