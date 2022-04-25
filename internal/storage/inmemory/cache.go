package inmemory

import "github.com/kapralovs/wb0/internal/orders"

type Cache map[string]*orders.Order
