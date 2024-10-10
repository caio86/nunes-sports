import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListagemProdutoComponent } from './listagem-produto/listagem-produto.component';
import { HttpClientModule } from '@angular/common/http';
import { MaterialModule } from 'src/app/shared/material/material.module';
import { FormProdutoComponent } from './form-produto/form-produto.component';
import { RouterModule } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';



@NgModule({
  declarations: [
    ListagemProdutoComponent,
    FormProdutoComponent,
  ],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MaterialModule,
  ],
  exports: [
    ListagemProdutoComponent,
    FormProdutoComponent,
  ]
})
export class ProdutoModule { }
