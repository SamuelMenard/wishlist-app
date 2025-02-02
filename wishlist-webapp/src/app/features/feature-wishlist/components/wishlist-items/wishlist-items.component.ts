import { ChangeDetectionStrategy, Component, inject, input } from '@angular/core';
import { WishlistItem } from '../../services/data/models/wishlist-item.model';
import { Router } from '@angular/router';

@Component({
    selector: 'app-wishlist-items',
    imports: [],
    templateUrl: './wishlist-items.component.html',
    styleUrl: './wishlist-items.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class WishlistItemsComponent {
    public readonly items = input.required<WishlistItem[]>();
    public readonly loading = input.required();
    public readonly error = input.required();

    private readonly router = inject(Router);

    protected onItemClicked(itemId: string): void {
        this.router.navigateByUrl(`wishlist/item/${itemId}`); // Navigate without reloading the page
    }
}
