import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { WishlistItem } from '../../services/data/models/wishlist-item.model';
import { catchError, finalize, of, take } from 'rxjs';
import { WishlistItemsClientService } from '../../services/data/clients/wishlist-items-client.service';

@Component({
    selector: 'app-page-wishlist-item',
    imports: [],
    templateUrl: './page-wishlist-item.component.html',
    styleUrl: './page-wishlist-item.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PageWishlistItemComponent implements OnInit {
    protected readonly model = signal<WishlistItem | null>(null);
    protected readonly isLoading = signal(true);
    protected readonly error = signal(false);
    private readonly route = inject(ActivatedRoute);
    private readonly wishlistItemsClientService = inject(WishlistItemsClientService);

    ngOnInit(): void {
        // Load selected model
        this.loadWishlistItem();
    }

    private loadWishlistItem(): void {
        const wishlistItemId = this.route.snapshot.paramMap.get('id');
        if (!wishlistItemId) {
            this.isLoading.set(false);
            return;
        }

        this.wishlistItemsClientService
            .getWishlistItemById(wishlistItemId)
            .pipe(
                catchError((err) => {
                    console.error('Error when loading wishlist item model', err);
                    return of(null);
                }),
                take(1),
                finalize(() => this.isLoading.set(false))
            )
            .subscribe((wishlistItemModel) => {
                // Failed to load model, return
                if (wishlistItemModel === null) {
                    return;
                }
                this.model.set(wishlistItemModel);
            });
    }
}
