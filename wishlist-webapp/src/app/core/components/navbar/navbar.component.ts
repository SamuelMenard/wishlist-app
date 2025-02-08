import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { WishlistClientService } from '../../../features/feature-wishlist/services/data/clients/wishlist-client.service';
import { Wishlist } from '../../../features/feature-wishlist/services/data/models/wishlist.model';
import { catchError, finalize, of, take } from 'rxjs';

@Component({
    selector: 'app-navbar',
    imports: [],
    templateUrl: './navbar.component.html',
    styleUrl: './navbar.component.scss',
    standalone: true,
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class NavbarComponent {
    protected readonly wishlists = signal<Wishlist[]>([]);
    protected readonly loading = signal(true);
    protected readonly error = signal(false);

    private readonly wishlistClientService = inject(WishlistClientService);

    constructor() {
        this.loadWishlists();
    }

    private loadWishlists(): void {
        this.wishlistClientService
            .getWishlists()
            .pipe(
                take(1),
                catchError((err) => {
                    console.log('Failed to load wishlists', err);
                    return of(null);
                }),
                finalize(() => this.loading.set(false))
            )
            .subscribe((wishlists) => {
                // Nothing to display, return
                if (!wishlists) {
                    return;
                }

                this.wishlists.set(wishlists);
            });
    }
}
