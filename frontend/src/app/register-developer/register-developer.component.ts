import { Component, OnInit } from '@angular/core';
import { DeveloperInput } from '../developer';
import { DeveloperService } from '../developer.service';

@Component({
  selector: 'app-register-developer',
  templateUrl: './register-developer.component.html',
  styleUrls: ['./register-developer.component.css'],
})
export class RegisterDeveloperComponent implements OnInit {
  constructor(private developerService: DeveloperService) {}

  ngOnInit(): void {}

  registerDeveloper(user_name: string, password: string): void {
    if (!user_name) {
      return;
    }
    const developerInput: DeveloperInput = {
      user_name: user_name,
      password: password,
    };
    this.developerService.addDeveloper(developerInput).subscribe();
  }
}
