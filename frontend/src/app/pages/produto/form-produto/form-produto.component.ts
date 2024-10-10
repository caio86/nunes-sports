import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Produto } from 'src/app/shared/models/Produto';
import { ProdutoService } from 'src/app/shared/services/produto.service';

@Component({
  selector: 'app-form-produto',
  templateUrl: './form-produto.component.html',
  styleUrls: ['./form-produto.component.scss']
})
export class FormProdutoComponent implements OnInit {
  formGroup: FormGroup
  isEditMode: boolean

  constructor(
    private produtoService: ProdutoService,
    private route: ActivatedRoute,
    private router: Router,
  ) {
    this.isEditMode = false
    this.formGroup = new FormGroup({
      id: new FormControl("", [Validators.required, Validators.pattern("^[0-9]+$")]),
      name: new FormControl("", Validators.required),
      description: new FormControl(null),
      price: new FormControl("", Validators.required),
    })
  }

  ngOnInit(): void {
    const id = Number(this.route.snapshot.paramMap.get("id"))
    if (id) {
      this.isEditMode = true
      this.produtoService.pesquisarPorID(id).subscribe({
        next: (value) => {
          this.formGroup.patchValue(value)
        },
        error: (err) => {
          console.error(err);
          alert("Produto não existe")
        },
      })
    }
  }

  getErrorMessage(field: string) {
    if (this.formGroup.get(field)?.hasError("required")) {
      return "Campo obrigatório"
    }

    if (this.formGroup.get(field)?.hasError("pattern")) {
      return `Valor inválido`
    }

    return "Alguma coisa deu errado"
  }

  cadastarProduto(produto: Produto): void {
    this.produtoService.inserir(produto).subscribe({
      next: () => {
        alert("Cliente cadastrado com sucesso!")
        this.router.navigate(["/produto"])
      },
      error: (err) => {
        console.error(err);
        alert("Erro ao cadastrar cliente!")
      }
    })
  }
  atualizarProduto(produto: Produto): void {
    this.produtoService.atualizar(produto).subscribe({
      next: () => {
        alert("Cliente atualizado com sucesso!")
        this.router.navigate(["/produto"])
      },
      error: (err) => {
        console.error(err);
        alert("Erro ao atualizar cliente!")
      }
    })
  }

  save(): void {
    const produto: Produto = this.formGroup.value
    if (!this.isEditMode) {
      this.cadastarProduto(produto)
    } else {
      this.atualizarProduto(produto)
    }
  }
}
