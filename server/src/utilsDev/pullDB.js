// деплой Redis
// ! заменить эту хуйню на нормальный shell скрипт

import redisClient from '@db/redis/index';

// import cityMap from '../map.json';

(async function pullCitiesDB() {
    try {
        // redisClient.set('ua', JSON.stringify(cityMap));
    } catch(err) {
        console.log('Error: ', err);
    }
})();
