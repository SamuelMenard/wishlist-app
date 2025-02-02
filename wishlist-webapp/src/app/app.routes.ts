import { Routes } from '@angular/router';
import { PageWishlistComponent } from './features/feature-wishlist/pages/page-wishlist/page-wishlist.component';
import { PageWishlistItemComponent } from './features/feature-wishlist/pages/page-wishlist-item/page-wishlist-item.component';

export const routes: Routes = [
    { path: 'wishlist/:id', component: PageWishlistComponent },
    { path: 'wishlist/item/:id', component: PageWishlistItemComponent },
];
