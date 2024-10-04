export interface Produto {
  id: string
  name: string
  description: string
  price: number
}

export interface GetProdutosResponseDTO {
  products: Produto[]
  _total: number
  _page: number
  _limit: number
}
