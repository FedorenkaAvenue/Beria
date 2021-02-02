import { Router } from 'express';
import { json } from 'body-parser';

 // use(json()).

import cityListController from '@controllers/cities';

const router: Router = Router();

router.get(['/city/:country', '/city'], cityListController); // ? найти лучше решение для роутинга

export default router;
