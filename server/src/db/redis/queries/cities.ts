import redisClient from '@db/redis/index';
import { ICities } from '@interfaces/City';

export default async function(country: string): Promise<ICities> {
    return new Promise((resolve, reject) => {
        redisClient.get(country, (err, cities) => err ? reject(err) : resolve(JSON.parse(cities as string)) );
    });
}
