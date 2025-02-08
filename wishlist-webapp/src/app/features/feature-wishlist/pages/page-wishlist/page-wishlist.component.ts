import { Component, inject, OnInit, signal } from '@angular/core';
import { WishlistClientService } from '../../services/data/clients/wishlist-client.service';
import { ActivatedRoute } from '@angular/router';
import { Wishlist } from '../../services/data/models/wishlist.model';
import { catchError, finalize, of, take } from 'rxjs';
import { WishlistItem } from '../../services/data/models/wishlist-item.model';
import { WishlistItemsComponent } from '../../components/wishlist-items/wishlist-items.component';

@Component({
    selector: 'app-page-wishlist',
    imports: [WishlistItemsComponent],
    templateUrl: './page-wishlist.component.html',
    styleUrl: './page-wishlist.component.scss',
})
export class PageWishlistComponent implements OnInit {
    private readonly route = inject(ActivatedRoute);
    private readonly wishlistClient = inject(WishlistClientService);

    protected readonly isModelLoading = signal(true);
    protected readonly model = signal<Wishlist | null>(null);

    protected readonly items = signal<WishlistItem[]>([]);
    protected readonly areItemsLoading = signal(true);
    protected readonly wishlistItemsLoadingError = signal(false);

    ngOnInit(): void {
        // Load selected model
        this.loadSelectedWishlist();
    }

    private loadSelectedWishlist(): void {
        const wishlistId = this.route.snapshot.paramMap.get('id');
        if (!wishlistId) {
            // Notify view that an error has occured
            this.isModelLoading.set(false);
            this.areItemsLoading.set(false);
            return;
        }

        // We want to load the model as fast as possible
        this.loadWishlistModel(wishlistId);

        // Items can take more time to load, load them separately
        this.loadWishlistItems(wishlistId);
    }

    private loadWishlistModel(wishlistId: string): void {
        this.wishlistClient
            .getWishlistById(wishlistId)
            .pipe(
                catchError((err) => {
                    console.error('Error when loading wishlist model', err);
                    return of(null);
                }),
                take(1),
                finalize(() => this.isModelLoading.set(false))
            )
            .subscribe((wishlistModel) => {
                // Failed to load model, return
                if (wishlistModel === null) {
                    return;
                }
                this.model.set(wishlistModel);
            });
    }

    private loadWishlistItems(wishlistId: string): void {
        this.wishlistClient
            .getWishlistItems(wishlistId)
            .pipe(
                catchError((err) => {
                    console.error('Error when loading wishlist items', err);
                    this.wishlistItemsLoadingError.set(true);
                    return of(null);
                }),
                take(1),
                finalize(() => this.areItemsLoading.set(false))
            )
            .subscribe((wishlistItems) => {
                // Failed to load wishlist items, return
                if (wishlistItems === null) {
                    return;
                }
                this.items.set(wishlistItems);
            });
    }
}
