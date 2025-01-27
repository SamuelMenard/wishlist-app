import { WishlistItemStatus } from '../enums/wishlist-item-status.enum';

export interface WishlistItem {
    id: string;
    name: string;
    status: WishlistItemStatus;
    createdAt: Date;
}
