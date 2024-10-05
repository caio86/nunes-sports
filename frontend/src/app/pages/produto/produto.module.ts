import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListagemProdutoComponent } from './listagem-produto/listagem-produto.component';
import { HttpClientModule } from '@angular/common/http';
import { MaterialModule } from 'src/app/shared/material/material.module';



@NgModule({
  declarations: [
    ListagemProdutoComponent
  ],
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
  ],
  exports: [
    ListagemProdutoComponent
  ]
})
export class ProdutoModule { }
