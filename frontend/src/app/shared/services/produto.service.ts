import { HttpClient } from '@angular/common/http'
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { GetProdutosResponseDTO, Produto } from '../models/Produto';

@Injectable({
  providedIn: 'root'
})
export class ProdutoService {
  api = `${environment.api}/product`

  constructor(
    private clienteHttp: HttpClient,
  ) { }

  listar(): Observable<GetProdutosResponseDTO> {
    return this.clienteHttp.get<GetProdutosResponseDTO>(this.api)
  }

  listarPaginado(page: number, pageSize: number): Observable<GetProdutosResponseDTO> {
    return this.clienteHttp.get<GetProdutosResponseDTO>(
      `${this.api}?page=${page}&limit=${pageSize}`
    )
  }

  inserir(novoProduto: Produto): Observable<Produto> {
    return this.clienteHttp.post<Produto>(
      this.api, novoProduto
    )
  }
}
