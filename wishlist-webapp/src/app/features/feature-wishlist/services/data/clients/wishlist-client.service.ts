import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Wishlist } from '../models/wishlist.model';
import { WishlistItem } from '../models/wishlist-item.model';

@Injectable({
    providedIn: 'root',
})
export class WishlistClientService {
    private readonly apiUrl = 'http://localhost:8080';
    private readonly http = inject(HttpClient);

    public getWishlistById(id: string): Observable<Wishlist> {
        return this.http.get<Wishlist>(`${this.apiUrl}/wishlist/${id}`);
    }

    public getWishlistItems(wishlistId: string): Observable<WishlistItem[]> {
        return this.http.get<WishlistItem[]>(`${this.apiUrl}/wishlist/items/${wishlistId}`);
    }

    public getWishlists(): Observable<Wishlist[]> {
        return this.http.get<Wishlist[]>(`${this.apiUrl}/wishlists`);
    }
}
