import { AfterViewInit, Component } from '@angular/core';
import { PageEvent } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { Produto } from 'src/app/shared/models/Produto';
import { ProdutoService } from 'src/app/shared/services/produto.service';

@Component({
  selector: 'app-listagem-produto',
  templateUrl: './listagem-produto.component.html',
  styleUrls: ['./listagem-produto.component.scss']
})
export class ListagemProdutoComponent implements AfterViewInit {
  displayedColumns: string[] = ["id", "name", "description", "price", "funcoes"]
  dataSource = new MatTableDataSource<Produto>
  dataSourceLength = 0
  page = 1
  pageSize = 5

  constructor(
    private produtoService: ProdutoService
  ) { }

  ngAfterViewInit(): void {
    this.listarProdutos(this.page, this.pageSize)
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
    this.page = event.pageIndex + 1
    this.pageSize = event.pageSize
    this.listarProdutos(this.page, this.pageSize)
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value
    this.dataSource.filter = filterValue.trim().toLowerCase()
  }

  deletarCliente(id: number) {
    if (confirm("Você tem certeza que deseja deletar?")) {
      this.produtoService.deletar(id).subscribe({
        next: () => {
          alert("Cliente deletado com sucesso!")
          this.listarProdutos(this.page, this.pageSize)
        },
        error: (err) => {
          console.error(err);
          alert("Erro ao deletar cliente!")
        }
      })
    }
  }
}
