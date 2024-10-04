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
}
