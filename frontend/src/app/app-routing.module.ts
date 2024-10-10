import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ListagemProdutoComponent } from './pages/produto/listagem-produto/listagem-produto.component';
import { FormProdutoComponent } from './pages/produto/form-produto/form-produto.component';

const routes: Routes = [
  {
    path: "produto",
    children: [
      {
        path: "novo",
        component: FormProdutoComponent,
      },
      {
        path: "editar/:id",
        component: FormProdutoComponent,
      },
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
