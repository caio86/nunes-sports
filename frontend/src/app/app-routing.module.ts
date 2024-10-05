import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ListagemProdutoComponent } from './pages/produto/listagem-produto/listagem-produto.component';

const routes: Routes = [
  {
    path: "produto",
    children: [
      {
        path: "",
        component: ListagemProdutoComponent,
      },
    ],
  },
  {
    path: "",
    component: ListagemProdutoComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
