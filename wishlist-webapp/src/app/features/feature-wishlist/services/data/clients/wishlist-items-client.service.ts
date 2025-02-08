import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { WishlistItem } from '../models/wishlist-item.model';

@Injectable({
    providedIn: 'root',
})
export class WishlistItemsClientService {
    private readonly apiUrl = 'http://localhost:8080';
    private readonly http = inject(HttpClient);

    public getWishlistItemById(wishlistItemId: string): Observable<WishlistItem> {
        return this.http.get<WishlistItem>(`${this.apiUrl}/wishlistItem/${wishlistItemId}`);
    }
}
