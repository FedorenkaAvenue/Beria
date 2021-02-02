export const { IS_DEV } = process.env; // ! определять перед импортом app

import './_moduleAliases';
import app from './app';

try {
    app.listen(process.env.APP_PORT || 80);
} catch(err) {
    console.error(`ошибка старта. ${err}`);
}
