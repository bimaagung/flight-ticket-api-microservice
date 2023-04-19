import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class User {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column('varchar', { length: '112', nullable: false })
  first_name: string;

  @Column('varchar', { length: '112', nullable: false })
  last_name: string;

  @Column({ type: 'int', width: 1, nullable: false })
  gender: number;

  @Column('text')
  address: string;

  @Column('varchar', { length: '112', unique: true, nullable: false })
  email: string;

  @Column('varchar', { length: '112', nullable: false })
  password: string;

  @Column({ type: 'boolean', default: false, nullable: false })
  admin: boolean;

  @Column({
    type: 'timestamp',
    default: () => 'CURRENT_TIMESTAMP',
    nullable: false,
  })
  created_at: string;

  @Column({
    type: 'timestamp',
    default: () => 'CURRENT_TIMESTAMP',
    nullable: false,
  })
  updated_at: string;

  @Column({ nullable: true })
  deleted_at: string;
}
