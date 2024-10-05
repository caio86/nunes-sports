import { Component, OnInit } from '@angular/core';
import { PageEvent } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { Produto } from 'src/app/shared/models/Produto';
import { ProdutoService } from 'src/app/shared/services/produto.service';

@Component({
  selector: 'app-listagem-produto',
  templateUrl: './listagem-produto.component.html',
  styleUrls: ['./listagem-produto.component.scss']
})
export class ListagemProdutoComponent implements OnInit {
  displayedColumns: string[] = ["id", "name", "description", "price"]
  dataSource = new MatTableDataSource<Produto>
  dataSourceLength: number = 0

  constructor(
    private produtoService: ProdutoService
  ) { }

  ngOnInit(): void {
    this.listarProdutos(1, 5)
  }

  listarProdutos(page: number, pageSize: number) {
    this.produtoService.listarPaginado(page, pageSize).subscribe(
      value => {
        this.dataSource.data = value.products
        this.dataSourceLength = value._total
      }
    )
  }

  handlePageEvent(event: PageEvent) {
    const page = event.pageIndex + 1
    const pageSize = event.pageSize
    this.listarProdutos(page, pageSize)
  }
}
